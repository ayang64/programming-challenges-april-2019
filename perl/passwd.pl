#!/usr/bin/perl
while (<>) {
    chop;
    $w[$#w + 1] = $_;
}
foreach $a (@w) {
    foreach $b (@w) {
        print "$a$b\n";
        print "$a $b\n";
    }
}
