# Design Pattern Examples with Golang

****************************************************************************************
# SOLID DESIGN PRINCIPLES

----------------------------------------REF---------------------------------------------

1. https://dave.cheney.net/2016/08/20/solid-go-design

----------------------------------------REF---------------------------------------------

# 1. Single Responsibility Principle

**     - A class should have one, and only one, reason to change.
                                                     -Robert C Martin

     - Seperation of concerns - different types/packages handling different, independent
       tasks/problems.

     - The Single Responsibility Principle encourages you to structure the functions, 
       types, and methods into packages that exhibit natural cohesion; the types
       belong together, the functions serve a single purpose.**

# 2. Open Closed Principle

**     - Software entities should be open for extension, but closed for modification.
                                                                     -Bertrand Meyer
     
     - The Open / Closed Principle encourages you to compose simple types into more
       complex ones using embedding.**

# 3. Liskov Substitution Principle

**     - Require no more, promise no less.
                                -Jim Weirich

     - Not Applicable for GOLANG as it is Typed Language

     - In a class based language, Liskov’s substitution principle is commonly
       interpreted as a specification for an abstract base class with various concrete
       subtypes. But Go does not have classes, or inheritance, so substitution cannot
       be implemented in terms of an abstract class hierarchy.**

# 4. Interface Segregation Principle

**     - Clients should not be forced to depend on methods they do not use.
                                                            -Robert C. Martin

     - Don't put too much into an interface; split into seperate interfaces.

     - A great rule of thumb for Go is accept interface, return structs.
                                                            -Jack Lindamood**

# 5. Dependency Inversion Principle

**     - High-level modules should not depend on low-level modules. Both should 
       depend on abstractions. Abstractions should not depend on details.
       Details should depend on abstractions.
                                                            -Robert C. Martin

     - The dependency inversion principle encourages you to push the responsibility
       for the specifics, as high as possible up the import graph, to your main
       package or top level handler, leaving the lower level code to deal with
       abstractions–interfaces.**

****************************************************************************************
