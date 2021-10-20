from abc import ABC,abstractmethod

class ClassWithInterface(ABC):
    @abstractmethod
    def say_hello(self):
        print("Hello")

# ClassWithInterface().say_hello()


# class NewClass(ClassWithInterface):
#     pass

# NewClass().say_hello()



# class NewClass(ClassWithInterface):
#     def say_hello(self):
#         print("No Hello!")

# NewClass().say_hello()