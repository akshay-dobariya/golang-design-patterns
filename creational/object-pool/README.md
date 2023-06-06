The Object Pool Design Pattern is a creational design pattern in which a pool of objects is initialized and created beforehand and kept in a pool. As and when needed, a client can request an object from the pool, use it, and return it to the pool. The object in the pool is never destroyed.

# When to Use:

- When the cost to create the object of the class is high and the number of such objects that will be needed at a particular time is not much.
- When the pool object is the immutable object.
- For performance reasons. It will boost the application performance significantly since the pool is already created
