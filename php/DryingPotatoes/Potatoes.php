<?php

function potatoes($p0, $w0, $p1): int {
    return (int) $w0 * (100 - $p0) / (100 - $p1);
}

print Potatoes(99,100,98); //50
print "\n";
print potatoes(82,127,80); // 114
