---
layout: post
title: "Python, Ruby, and Golang: A Web Service Application Comparison"
author: Kyle W. Purdon
categories: [python, ruby, golang, development]
---

INTRO HERE

<!-- more -->

## Service Description

The service created provides a very basic blog. The following routes are constructed:

* `GET /`: Return the blog (using a template to render).
* `GET /json`: Return the blog content in JSON format.
* `POST /new`: Add a new post (title, summary, content) to the blog.

The external interface to the blog service is exactly the same for each language.

**Add A Post**

`POST /new`

{% highlight bash %}
curl --form title='Test Post 1' \
     --form summary='The First Test Post' \
     --form content='Lorem ipsum dolor sit amet, consectetur ...' \
     http://[IP]:[PORT]/new
 {% endhighlight %}

**View The HTML**

`GET /`

![blog preview png]({{root_url}}/images/blog.png)

**View The JSON**

`GET /json`

{% highlight js %}
[
   {
      content:"Lorem ipsum dolor sit amet, consectetur ...",
      title:"Test Post 1",
      _id:{
         $oid:"558329927315660001550970"
      },
      summary:"The First Test Post"
   }
]
{% endhighlight %}

## Application Structure

Each application can be broken down into the following components:

**Application Setup**

* Initialize an application
* Run the application

**Request**

* Define routes on which a user can request data (GET)
* Define routes on which a user can submit data (POST)

**Response**

* Render JSON (`GET /json`)
* Render a template (`GET /`)

**Database**

* Initialize a connection
* Insert data
* Retrieve data

**Application Deployment**

* Docker!

The rest of this article will compare each of these components for each library. The purpose is not to suggest that one of these libraries is better than the other rather it is to provide a specific comparison between the three tools:

* [flask](http://flask.pocoo.org/) (Python)
* [sinatra](http://www.sinatrarb.com/) (Ruby)
* [martini](http://martini.codegangsta.io/) (Golang)

## Initialize/Run An Application

**Python (flask)**

{% highlight python %}
# initialize application
from flask import Flask
app = Flask(__name__)

# run application
if __name__ == '__main__':
    app.run(host='0.0.0.0')
{% endhighlight %}

{% highlight bash %}
$ python app.py
{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
# initialize application
require 'sinatra'
{% endhighlight %}

{% highlight bash %}
$ ruby app.rb
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
// initialize application
import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
app := martini.Classic()
app.Use(render.Renderer())

// run application
app.Run()
{% endhighlight %}

{% highlight bash %}
$ go run app.go
{% endhighlight %}

## Define A Route (GET/POST)

**Python (flask)**

{% highlight python %}
# get
@app.route('/')
def blog():
    ...

#post
@app.route('/new', methods=['POST'])
def new():
    ...

{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
# get
get '/' do
  ...
end

# post
post '/new' do
  ...
end
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
// define data struct
type Post struct {
	Title   string `form:"title" json:"title"`
	Summary string `form:"summary" json:"summary"`
	Content string `form:"content" json:"content"`
}

// get
app.Get("/", func(r render.Render) {
	...
}

// post
import "github.com/martini-contrib/binding"
app.Post("/new", binding.Bind(Post{}), func(r render.Render, post Post) {
	...
}
{% endhighlight %}

## Render A JSON Response

**Python (flask)**

Flask provides a [jsonify]() method but since the service is using mongodb the mongodb bson utility is used.
{% highlight python %}
from bson.json_util import dumps
return dumps(posts) # posts is a list of dicts [{}, {}]
{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
require 'json'
content_type :json
posts.to_json # posts is an array (from mongodb)
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
r.JSON(200, posts) // posts is an array of Post{} structs
{% endhighlight %}

## Render An HTML Response (Templating)

**Python (flask)**

{% highlight python %}
return render_template('blog.html', posts=posts)
{% endhighlight %}

blog.html liquid template
{% highlight html %}
<!doctype HTML>
<html>
  <head>
    <title>Python Flask Example</title>
  </head>
  <body>{% raw %}
    {% for post in posts %}
      <h1> {{ post.title }} </h1>
      <h3> {{ post.summary }} </h3>
      <p> {{ post.content }} </p>
      <hr>
    {% endfor %}
  {% endraw %}</body>
</html>
{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
erb :blog
{% endhighlight %}

blog.erb erb template
{% highlight erb %}
<!doctype HTML>
<html>
  <head>
    <title>Ruby Sinatra Example</title>
  </head>
  <body>
    <% @posts.each do |post| %>
      <h1><%= post['title'] %></h1>
      <h3><%= post['summary'] %></h3>
      <p><%= post['content'] %></p>
      <hr>
    <% end %>
  </body>
</html>
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
r.HTML(200, "blog", posts)
{% endhighlight %}

blog.tmpl go template
{% highlight html %}
<!doctype HTML>
<html>
  <head>
    <title>Golang Martini Example</title>
  </head>
  <body>{% raw %}
    {{range . }}
      <h1>{{.Title}}</h1>
      <h3>{{.Summary}}</h3>
      <p>{{.Content}}</p>
      <hr>
    {{ end }}
  {% endraw %}</body>
</html>
{% endhighlight %}

## Database Connection

All of the applications are using the mongodb driver. The environment variable `DB_PORT_27017_TCP_ADDR` is the IP of a linked docker container (the database ip).

**Python (flask)**

{% highlight python %}
from pymongo import MongoClient
client = MongoClient(os.environ['DB_PORT_27017_TCP_ADDR'], 27017)
db = client.blog
{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
require 'mongo'
db_ip = [ENV['DB_PORT_27017_TCP_ADDR']]
client = Mongo::Client.new(db_ip, database: 'blog')
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
import "gopkg.in/mgo.v2"
session, _ := mgo.Dial(os.Getenv("DB_PORT_27017_TCP_ADDR"))
db := session.DB("blog")
defer session.Close()
{% endhighlight %}

## Insert Data From a POST

**Python (flask)**

{% highlight python %}
from flask import request
post = {
    'title': request.form['title'],
    'summary': request.form['summary'],
    'content': request.form['content']
}
db.blog.insert_one(post)
{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
client[:posts].insert_one(params) # params is a hash generated by sinatra
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
db.C("posts").Insert(post) // post is an instance of the Post{} struct
{% endhighlight %}

## Retrieve Data

**Python (flask)**

{% highlight python %}
_posts = db.blog.find()
posts = [post for post in _posts]
{% endhighlight %}

**Ruby (sinatra)**

{% highlight ruby %}
@posts = client[:posts].find.to_a
{% endhighlight %}

**Golang (martini)**

{% highlight go %}
var posts []Post
db.C("posts").Find(nil).All(&posts)
{% endhighlight %}

## Application Deployment (Docker!)

A great solution to deploying all of these applications is to use [Docker](https://www.docker.com/) and [Docker-Compose](https://docs.docker.com/compose/).

**Python (flask)**

*Dockerfile*
{% highlight docker %}
FROM python:2.7

ADD . /app
WORKDIR /app

RUN pip install -r requirements.txt
{% endhighlight %}

*docker-compose.yml*
{% highlight yaml %}
web:
  build: .
  command: python -u app.py
  ports:
    - "5000:5000"
  volumes:
    - .:/app
  links:
    - db
db:
  image: mongo:3.0.4
  command: mongod --quiet --logpath=/dev/null
{% endhighlight %}

**Ruby (sinatra)**

*Dockerfile*
{% highlight docker %}
FROM ruby:2.2

ADD . /app
WORKDIR /app

RUN bundle install
{% endhighlight %}

*docker-compose.yml*
{% highlight yaml %}
web:
  build: .
  command: bundle exec ruby app.rb
  ports:
    - "4567:4567"
  volumes:
    - .:/app
  links:
    - db
db:
  image: mongo:3.0.4
  command: mongod --quiet --logpath=/dev/null
{% endhighlight %}

**Golang (martini)**

*Dockerfile*
{% highlight docker %}
FROM golang:1.3

ADD . /go/src/github.com/kpurdon/go-todo
WORKDIR /go/src/github.com/kpurdon/go-todo

RUN go get github.com/go-martini/martini && go get github.com/martini-contrib/render && go get gopkg.in/mgo.v2 && go get github.com/martini-contrib/binding
{% endhighlight %}

*docker-compose.yml*
{% highlight yaml %}
web:
  build: .
  command: go run app.go
  ports:
    - "3000:3000"
  volumes: # look into volumes v. "ADD"
    - .:/go/src/github.com/kpurdon/go-todo
  links:
    - db
db:
  image: mongo:3.0.4
  command: mongod --quiet --logpath=/dev/null
{% endhighlight %}

## Conclusion

Docker feels like a heavy choice for this type of application unless high performance is a key requirement. Ruby and Python are likely going to be the better choice of language. Python and Ruby (flask and sintra) are very similar as far as this application goes. Here are a few notable differences:

### Simplicity

While flask is very lightweight and reads clearly the sinatra app is the simplest of the two. At 23 LOC (compared to 46 for flask and 42 for martini). Also the handling of input forms is handled behind the scenes in sinatra. For these reasons sinatra is the winner in this catagory.

### Documentation

The flask documentation was the simplest to search and most approachable. While the sinatra and martini documentation is complete it was not as approachable. For this reason flask is the winner in this catagory.

## Final Determination

The correct tool for this example is a tie between Python and Ruby. Pick whichever you are more comfortable with and you will be succesfull. If you need high performance consider Golang.
