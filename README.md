netrunner-limited
=================

Code for generating netrunner-limited formats. 

Currently supporting sealed (similar to MTG) where you start with a fixed pool 
of cards.

Will eventually support external flags.  Right now, the important constants to
modify are:

CardsPerDeck (controls the size of the generate pool)
RandSeed (any number would do. If you want to generate the same pool again, 
remember that number!)



