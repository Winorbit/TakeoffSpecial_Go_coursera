from datetime import datetime
now = datetime.now()

# НАСЛЕДОВАНИЕ
class BaseTextModel:
	title = "test"
	text = "test"
	date = now.strftime("%d/%m/%Y %H:%M:%S")

	def add_text(self, text):
		self.text = text

	def add_title(self, title):
		self.title = title

"""
Хоть BlogPost и пустой, но т.к он отнаследован от BaseTextModel,
он обладает всеми атрибутами и методами ,что есть у BaseTextModel.
"""
class BlogPost(BaseTextModel):
	pass

test_empty_post = BlogPost()
print(test_empty_post.title)
print(test_empty_post.text)
print(test_empty_post.date)


post_about_cowboy_beebop = BlogPost()
post_about_cowboy_beebop.add_text("один из множеста текстов про то, какое выдающееся аниме 'Ковбой Бибоп'")
post_about_cowboy_beebop.add_title("Real Folk Blues")

print(post_about_cowboy_beebop.title)
print(post_about_cowboy_beebop.text)
print(post_about_cowboy_beebop.date)


"""
Когда у класса-наследника объявлены методы и атрибуты
Они прибавляются к методам и атрибутам родителя. 
"""
class BlogPost(BaseTextModel):
	author = ""

	def add_author(self,author_username):
		self.author = author_username

	def post_as_dict(self):
		post = {}
		post["date"] = self.date
		post["title"] = self.title
		post["text"] = self.text
		post["author"] = self.author
		return post

post_about_cowboy_beebop = BlogPost()
post_about_cowboy_beebop.add_text("один из множеста текстов про то, какое выдающееся аниме 'Ковбой Бибоп'")
post_about_cowboy_beebop.add_title("Real Folk Blues")

post_as_dict = post_about_cowboy_beebop.post_as_dict()
print(post_as_dict)