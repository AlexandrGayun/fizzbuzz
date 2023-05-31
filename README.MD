# Main Idea
As FizzBuzz is just a set of consecutive checks, we could represent it as a some sort of pipeline that working with data step by step.
So in purpose to make this process as much flexible and extensible as possible we could provide a tool to build a such kind of pipelines.
It allows you to made your own pipelines with the steps you define that will execute in order and with rules provided by your own.
 
## How to use
The concept is super easy. The pipeline(processor) is a wrapper that provides a basic type operations(CreateNewProcessor/Run/etc what do you need). They are covered by the 
Core interface methods(WriteHandlers and Process, could be extended). Core is the structure which define a pipeline by its own.

## Example
Basically fizzbuzz is a consecutive conveyor which perform a three consecutive checks(fizzbuzz/fizz/buzz).
This kind of pipeline is implemented in linear package. We provide a Core struct which represent a set of operations(in WriteHandlers it just store them in a way of slice of handlers functions).
And in the Process method we were just applying them continuously one by one.

## Flexibility/Usability
We could extend it with any new handlers(steps) without rewriting previous implementations, reuse handlers in a different orders, reuse the cores we writed before, change the algorithms of handlers processing and much more