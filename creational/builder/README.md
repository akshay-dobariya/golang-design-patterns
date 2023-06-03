Builder Pattern is a creational design pattern used for constructing complex objects.

# When To Use

- Use Builder pattern when the object constructed is big and requires multiple steps. It helps in less size of the constructor. The construction of the computer becomes simple and it does not require a large constructor
- When a different version of the same product needs to be created. For example, in the below code we see a different version of computer
- When half constructed final object should not exist. Again referring to below code the computer created will either be created fully or not created at all. The Concrete Builder struct holds the temporary state of house object being created
