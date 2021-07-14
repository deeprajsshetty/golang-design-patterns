# Design Pattern Examples with Golang

****************************************************************************************
# SOLID DESIGN PRINCIPLES

----------------------------------------REF---------------------------------------------

1. https://dave.cheney.net/2016/08/20/solid-go-design

----------------------------------------REF---------------------------------------------

# 1. Single Responsibility Principle

     - A class should have one, and only one, reason to change.
                                                     Robert C Martin

     - Seperation of concerns - different types/packages handling different, independent
       tasks/problems.

     - The Single Responsibility Principle encourages you to structure the functions, 
       types, and methods into packages that exhibit natural cohesion; the types
       belong together, the functions serve a single purpose.

# 2. Open Closed Principle

     - Software entities should be open for extension, but closed for modification.
                                                      Bertrand Meyer
     
     - The Open / Closed Principle encourages you to compose simple types into more
       complex ones using embedding.

# 3. Liskov Substitution Principle

     - Not Applicable for GOLANG as it is Typed Language

# 4. Interface Segregation Principle

     - Don't put too much into an interface; split into seperate interfaces

# 5. Dependency Inversion Principle

     - High-level modules should not depend upon low-level ones; use abstractions

****************************************************************************************