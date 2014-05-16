netrunner-limited
=================

Code for generating netrunner-limited formats. 

Currently supporting sealed (similar to MTG) where you start with a fixed pool 
of cards.

1. CardsPerDeck (controls the size of the generate pool)
2. RandSeed (any number would do. If you want to generate the same pool again, 
remember that number!)

Note that golang map keys are random - you are generating the same pool though.

Example with external flags:

go run sealed_pool_creator.go -cards_per_deck 75 -random_seed 123456778

