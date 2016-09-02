namespace php tt.demo
namespace go lib.thrift.demo

enum Operation {
    ADD,
    SUB,
    MUL,
    DIV,
}

struct CalculateInput {
    1: required double number1;
    2: required double number2;
    3: required Operation op;
}

exception InputException {
    1: required string message,
}

// Thrift will generate client, so an engineer only has to implement server code and some
// bootstrapping for both client and server
service DemoService {
    // Calculate calculates outcome of desired operations on two numbers
    double calculate(1: CalculateInput input) throws (
        1:InputException invalidInput, // thrown when operation can't be completed
    ),

    // RandomGenerator returns _size_ numbers as list, in range from _min_ to _max_.
    list<i64> randomGenerator(1: required i64 min, 2: required i64 max, 3: required i64 size),
}
