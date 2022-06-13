package actions

import (
	"fmt"
	"net/http"

	"go_buffalo_api/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Book)
// DB Table: Plural (books)
// Resource: Plural (Books)
// Path: Plural (/books)
// View Template Folder: Plural (/templates/books/)

// BooksResource is the resource for the Book model
type BooksResource struct {
	buffalo.Resource
}

// List gets all Books. This function is mapped to the path
// GET /books
func (v BooksResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	books := &models.Books{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Books from the DB
	if err := q.All(books); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("books", books)
		return c.Render(http.StatusOK, r.HTML("books/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(books))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(books))
	}).Respond(c)
}

// Show gets the data for one Book. This function is mapped to
// the path GET /books/{book_id}
func (v BooksResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Book
	book := &models.Book{}

	// To find the Book the parameter book_id is used.
	if err := tx.Find(book, c.Param("book_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("book", book)

		return c.Render(http.StatusOK, r.HTML("books/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(book))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(book))
	}).Respond(c)
}

// Create adds a Book to the DB. This function is mapped to the
// path POST /books
func (v BooksResource) Create(c buffalo.Context) error {
	// Allocate an empty Book
	book := &models.Book{}

	// Bind book to the html form elements
	if err := c.Bind(book); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(book)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("book", book)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("books/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "book.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/books/%v", book.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(book))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(book))
	}).Respond(c)
}

// Update changes a Book in the DB. This function is mapped to
// the path PUT /books/{book_id}
func (v BooksResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Book
	book := &models.Book{}

	if err := tx.Find(book, c.Param("book_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Book to the html form elements
	if err := c.Bind(book); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(book)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("book", book)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("books/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "book.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/books/%v", book.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(book))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(book))
	}).Respond(c)
}

// Destroy deletes a Book from the DB. This function is mapped
// to the path DELETE /books/{book_id}
func (v BooksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Book
	book := &models.Book{}

	// To find the Book the parameter book_id is used.
	if err := tx.Find(book, c.Param("book_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(book); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "book.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/books")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(book))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(book))
	}).Respond(c)
}
