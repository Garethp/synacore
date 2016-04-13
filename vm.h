#ifndef SYNACORE_VM_H
#define SYNACORE_VM_H

#include <iostream>
#include "memory.h"
#include "registers.h"
#include <stdint.h>

using namespace std;

class vm {
    memory memory1;
    registers register1;

    uint32_t intLimit = 32768;
    bool halted = false;


    void doOperation(uint32_t operation);

    void halt();
    void setRegistry();
    void isEqual();
    void greaterThan();
    void jumpTo();
    void jt();
    void jf();
    void add();
    void multiply();
    void modulo();
    void print();
    void noop();

    uint32_t getValue(uint32_t value);
    uint32_t getFromRegister(uint32_t index);
    void setToRegister(uint32_t index, uint32_t value);
    bool shouldKeepRunning();
public:
    void run();
};


#endif //SYNACORE_VM_H
