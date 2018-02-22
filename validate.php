<?php

$r = [4,1,0,0,0,0,0,3];
$stack = [];
$mem = [];
$lastStack = null;

//6027
function validate() {
    global $r;
    global $stack;
    global $mem;

    //6035
    if ($r[0] !== 0) {
        //6048
        if ($r[1] !== 0) {
            $stack[] = $r[0];
//            $lastStack = $r[0];
//            $r[1] = add($r[1], 32767);
            $r[1] = 0;
            $r[0]--;
//            validate();

            $r[1] = $r[0];
//            $r[0] = $lastStack;
            $r[0] = array_pop($stack);
            $r[0] = add($r[0], 32767);
            validate();
            return;
        }

        //R0 --?
        $r[0] = add($r[0], 32767);
        $r[1] = $r[7];
//
        validate();
        return;
    }

    $r[0] = add($r[1], 1);

    return;
}

function add($a, $b) {
    $mod = 32768;

    return ($a + $b) % $mod;
}

//5498
$r[0] = 4;
$r[1] = 1;
validate();

//Set Register 0 to 6 and skip validation?!?