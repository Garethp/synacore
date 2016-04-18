#include "vm.h"
#include <iostream>
#include <stdint.h>

using namespace std;

bool vm::shouldKeepRunning() {
    return !this->halted && !this->memory1.isEOM();
}

void vm::run() {
    this->memory1.loadFromFile("challenge.bin");

    while (this->shouldKeepRunning()) {
        uint32_t operation = memory1.read();
        operation = getValue(operation);
        doOperation(operation);
    }
}

void vm::doOperation(uint32_t operation) {
    switch (operation) {
        case 0:
            this->halt();
            break;
        case 1:
            this->setRegistry();
            break;
        case 2:
            this->push();
            break;
        case 3:
            this->pop();
            break;
        case 4:
            this->isEqual();
            break;
        case 5:
            this->greaterThan();
            break;
        case 6:
            this->jumpTo();
            break;
        case 7:
            this->jumpToIfNotZero();
            break;
        case 8:
            this->jumpToIfZero();
            break;
        case 9:
            this->add();
            break;
        case 10:
            this->multiply();
            break;
        case 11:
            this->modulo();
            break;
        case 12:
            this->bitwiseAnd();
            break;
        case 13:
            this->bitwiseOr();
            break;
        case 14:
            this->bitwiseNot();
            break;
        case 15:
            this->readFromMemory();
            break;
        case 16:
            this->writeToMemory();
            break;
        case 17:
            this->call();
            break;
        case 18:
            this->retrieve();
            break;
        case 19:
            this->print();
            break;
        case 20:
            this->readCharacter();
            break;
        case 21:
            this->noop();
            break;
        default:
            cout << "Unknown Operation " << operation << endl;
            this->halted = true;
            break;
    }
}

void vm::halt() {
    this->halted = true;
}

void vm::push() {
    uint32_t value = getValue(memory1.read());
    this->stack1.push(value);
}

void vm::pop() {
    uint32_t registry = memory1.read();
    this->setToRegister(registry, this->stack1.top());
    this->stack1.pop();
}

void vm::setRegistry() {
    uint32_t registryIndex = memory1.read();
    uint32_t value = this->getValue(memory1.read());

    this->setToRegister(registryIndex, value);
}

void vm::isEqual() {
    uint32_t registryIndex = memory1.read();
    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    if (a == b) {
        setToRegister(registryIndex, 1);
    } else {
        setToRegister(registryIndex, 0);
    }
}

void vm::greaterThan() {
    uint32_t registryIndex = memory1.read();
    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    if (a > b) {
        setToRegister(registryIndex, 1);
    } else {
        setToRegister(registryIndex, 0);
    }
}

void vm::jumpTo() {
    uint32_t nextMemoryAddress = memory1.read();
    memory1.setPointer(nextMemoryAddress);
}

void vm::jumpToIfNotZero() {
    uint32_t value = memory1.read();
    value = getValue(value);
    uint32_t jumpValue = memory1.read();

    if (getValue(value) != 0) {
        this->memory1.setPointer(jumpValue);
    }
}

void vm::jumpToIfZero() {
    uint32_t value = getValue(memory1.read());
    uint32_t jumpValue = memory1.read();

    if (getValue(value) == 0) {
        this->memory1.setPointer(jumpValue);
    }
}

void vm::add() {
    uint32_t registerIndex = memory1.read();

    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    uint32_t c = (a + b) % this->intLimit;

    this->setToRegister(registerIndex, c);
}

void vm::multiply() {
    uint32_t registryIndex = memory1.read();

    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    uint32_t c = (a * b) % this->intLimit;

    this->setToRegister(registryIndex, c);
}

void vm::modulo() {
    uint32_t registryIndex = memory1.read();

    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    uint32_t c = a % b;

    this->setToRegister(registryIndex, c);
}

void vm::bitwiseAnd() {
    uint32_t registryIndex = memory1.read();
    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    uint32_t c = a & b;

    setToRegister(registryIndex, c);
}

void vm::bitwiseOr() {
    uint32_t registryIndex = memory1.read();
    uint32_t a = getValue(memory1.read());
    uint32_t b = getValue(memory1.read());

    uint32_t c = a | b;

    setToRegister(registryIndex, c);
}

void vm::bitwiseNot() {
    uint32_t registryIndex = memory1.read();
    uint32_t a = getValue(memory1.read());

    uint32_t c = (~a) % this->intLimit;

    setToRegister(registryIndex, c);
}

void vm::readFromMemory() {
    uint32_t registryIndex = memory1.read();
    uint32_t memoryAddress = getValue(memory1.read());

    setToRegister(registryIndex, memory1.read(memoryAddress));
}

void vm::writeToMemory() {
    uint32_t memoryAddress = getValue(memory1.read());
    uint32_t value = getValue(memory1.read());

    memory1.write(memoryAddress, value);
}

void vm::call() {
    uint32_t jumpTo = this->getValue(memory1.read());
    uint32_t nextInstruction = this->memory1.getPointer();

    this->stack1.push(nextInstruction);

    this->memory1.setPointer(jumpTo);
}

void vm::retrieve() {
    uint32_t jumpTo = this->stack1.top();
    this->stack1.pop();

    this->memory1.setPointer(jumpTo);
}

void vm::print() {
    uint32_t value = this->getValue(memory1.read());

    char character = (char) value;
    cout << character;
}

void vm::noop() {

}

uint32_t vm::getValue(uint32_t value) {
    if (value < this->intLimit) {
        return value;
    }

    //Return register value here
    return this->getFromRegister(value);
}

uint32_t vm::getFromRegister(uint32_t index) {
    index = index % this->intLimit;

    return this->register1.getRegister(index);
}

void vm::readCharacter() {
    uint32_t registryIndex = memory1.read();

    if (this->inputBuffer.length() == 0) {
        cin >> this->inputBuffer;
        inputBuffer += '\n';
    }

    char c = inputBuffer.at(0);
    inputBuffer.erase(0, 1);

    setToRegister(registryIndex, (uint32_t) c);
}

void vm::setToRegister(uint32_t index, uint32_t value) {
    index = index % this->intLimit;

    this->register1.setRegsiter(index, value);
}