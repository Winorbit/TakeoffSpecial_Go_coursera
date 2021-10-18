"""
Опишем с помощью классов такую сущность, как пост в блоге.
Зададим атрибуты, которые будут содержать данные нашего поста и
зададим два метода - функции, привязанные к конкретному классу.

add_text - позволяет изменить значение атрибута text
add_title - позволяет изменить значение атрибута text

Т.к изначально значения text и title пусты, мы будем менять его с помощью методов.
"""



class BlogPost:
	title = ""
	text = ""
	author = ""
	
	def add_text(self, text):
		"""
		self - обязательный аргумент, который являеться ссылкой объекта на самого себя.
		т.е self.text, получается - это переменная, привязанная не к классу а к объекту
		порожденному из этого класса.
		Сделав 10 объектов Lesson у этих 10 объектов будет 10 значений text равных "",
		но это будут 10 разных пустых строк. И self указывает, какой именно это text
		"""
		self.text = text

	def add_title(self, title):
		self.title = title
# Из класса BlogPost создаем объект BlogPost()
post_about_pandas = BlogPost()

# Меняем значения его атрибутов с помощью его же методов.
post_about_pandas.add_text("Любой текст поста про панд")
post_about_pandas.add_title("Просто заголовок")


print(post_about_pandas.title)
print(post_about_pandas.text)
print(post_about_pandas.date)


post_about_cowboy_beebop = BlogPost()
post_about_cowboy_beebop.add_text("Один из множеста текстов про то, какое выдающееся аниме 'Ковбой Бибоп'")
post_about_cowboy_beebop.add_title("Real Folk Blues")

print(post_about_cowboy_beebop.title)
print(post_about_cowboy_beebop.text)
print(post_about_cowboy_beebop.date)