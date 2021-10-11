lambda x:print(x)
print(lambda x:print(x))

(lambda x:print(x))("Hello")


for x in (1,"Hello",None,4.5):
	(lambda arg:print(f"Type of argument is: {type(arg)}"))(x)

# Здесь лямбда-функция применяется для того, чтобы, отсортировать каждый элемент в списке lst, 
# т.е анонимная функция примменяеться на каждый из элментов.
lst = [('candy','30','100'), ('apple','10','200'), ('baby','20','300')]
lst.sort(key=lambda x:x[1])
print(lst)