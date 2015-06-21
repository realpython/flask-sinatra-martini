import os

from flask import Flask, render_template, request
from pymongo import MongoClient
from bson.json_util import dumps

app = Flask(__name__)

client = MongoClient(
    os.environ['DB_1_PORT_27017_TCP_ADDR'],
    27017)
db = client.blog


@app.route('/')
def blog():

    _posts = db.blog.find()
    posts = [post for post in _posts]

    return render_template('blog.html', posts=posts)


@app.route('/json')
def blog_json():

    _posts = db.blog.find()
    posts = [post for post in _posts]

    return dumps(posts)


@app.route('/new', methods=['POST'])
def new():

    post = {
        'title': request.form['title'],
        'summary': request.form['summary'],
        'content': request.form['content']
    }
    db.blog.insert_one(post)

    return 'Insert Complete'

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)
