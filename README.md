# Welcome to Buffalo

Head over to https://gobuffalo.io/ to view the full documentation.

## Database Setup

Start your postgres Docker container:

```console
docker-compose up
```

### Create Your Databases

Buffalo uses the "database.yml" file to know how to connect to your database. Once you've edited this file and started your database, now Buffalo can create the databases in that file for you:

```console
buffalo pop create -a
```

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```console
buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

The following steps will show some examples of how Buffalo makes your life easier. The only requirement is that you first install the [Soda CLI](https://gobuffalo.io/documentation/database/soda/).

## Generate a Model

[Documentation](https://gobuffalo.io/documentation/database/models/#using-the-generator)

This will generate the model files and the migrations to create the corresponding database tables. 

*Note: Currently generating associations (aka relationships) through the CLI is not supported. So you will have to generate the models and edit the resulting model and migration files manually. This is something on the Buffalo devs' radar.*
```console
soda generate model user title:string first_name:string last_name:string bio:text
```

## Generate Resources

[Documentation](https://gobuffalo.io/documentation/request_handling/resources/)

A "Resource" is a Buffalo's concept of Auto-CRUD. It speeds up development by auto-defining handlers and routes for CRUD operations.

```console
buffalo g resource books --use-model book --skip-templates
```
