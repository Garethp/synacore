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
        uint32_t operation = getValue(memory1.getCurrentMemory());
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
            this->jt();
            break;
        case 8:
            this->jf();
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
        case 19:
            this->print();
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

void vm::setRegistry() {
    uint32_t registryIndex = memory1.getCurrentMemory();
    uint32_t value = this->getValue(memory1.getCurrentMemory());

    this->setToRegister(registryIndex, value);
}

void vm::isEqual() {
    uint32_t registryIndex = memory1.getCurrentMemory();
    uint32_t a = getValue(memory1.getCurrentMemory());
    uint32_t b = getValue(memory1.getCurrentMemory());

    if (a == b) {
        setToRegister(registryIndex, 1);
    } else {
        setToRegister(registryIndex, 0);
    }
}

void vm::greaterThan() {
    uint32_t registryIndex = memory1.getCurrentMemory();
    uint32_t a = getValue(memory1.getCurrentMemory());
    uint32_t b = getValue(memory1.getCurrentMemory());

    if (a > b) {
        setToRegister(registryIndex, 1);
    } else {
        setToRegister(registryIndex, 0);
    }
}

void vm::jumpTo() {
    uint32_t nextMemoryAddress = memory1.getCurrentMemory();
    memory1.setPointer(nextMemoryAddress);
}

void vm::jt() {
    uint32_t value = memory1.getCurrentMemory();
    value = getValue(value);
    uint32_t jumpValue = memory1.getCurrentMemory();

    if (getValue(value) != 0) {
        this->memory1.setPointer(jumpValue);
    }
}

void vm::jf() {
    uint32_t value = getValue(memory1.getCurrentMemory());
    uint32_t jumpValue = memory1.getCurrentMemory();

    if (getValue(value) == 0) {
        this->memory1.setPointer(jumpValue);
    }
}

void vm::add() {
    uint32_t registerIndex = memory1.getCurrentMemory();

    uint32_t a = getValue(memory1.getCurrentMemory());
    uint32_t b = getValue(memory1.getCurrentMemory());

    uint32_t c = (a + b) % this->intLimit;

    this->setToRegister(registerIndex, c);
}

void vm::multiply() {
    uint32_t registryIndex = memory1.getCurrentMemory();

    uint32_t a = getValue(memory1.getCurrentMemory());
    uint32_t b = getValue(memory1.getCurrentMemory());

    uint32_t c = (a * b) % this->intLimit;

    this->setToRegister(registryIndex, c);
}

void vm::modulo() {
    uint32_t registryIndex = memory1.getCurrentMemory();

    uint32_t a = getValue(memory1.getCurrentMemory());
    uint32_t b = getValue(memory1.getCurrentMemory());

    uint32_t c = a % b;

    this->setToRegister(registryIndex, c);
}

void vm::print() {
    uint32_t value = this->getValue(memory1.getCurrentMemory());

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

void vm::setToRegister(uint32_t index, uint32_t value) {
    index = index % this->intLimit;

    this->register1.setRegsiter(index, value);
}