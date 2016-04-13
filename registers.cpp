#include "registers.h"

#include <stdint.h>
#include <stdint.h>

uint32_t registers::getRegister(uint32_t index) {
    return this->registers[index];
}

void registers::setRegsiter(uint32_t index, uint32_t value) {
    this->registers[index] = value;
}