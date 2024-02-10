# Cards

Adapted from the Udemy GoLang masterclass first project.

This project will take the skills learned in the first unit as well as my previous experience in
development to become a full featured card game.

I will use https://ebitengine.org/ for the engine once I am ready to move the game from the
command line.

## Stage one

Early in the dev cycle i think its important to figure out the big hurdles that will come with 
more advanced features. To this end i will implement Texas Hold'em as my first game.

### Texas Hold'em

this game will have all the features that will be necessary for later games, i believe.

For example;
    menu
    Player hands
    Dealer hands/community cards
    turns-taking
    basic ai :-/
    chips/betting
    win/loss hierarchy

menu -> turns taking -> hands -> everything else

making a menu will allow me to create a minimum viable product as quickly as possible
offering the ability to test as i build and offering a framework to build gameplay around.

chronology should naturally be among the foundational features to ensure i have a good base to build everything else on.

#### menu

menu function

    func menu(string label,[]string items) string {
        _, result, err := select{
            Label:
            Items:
        }

        if err != nil {
            fmt.Println("Menu Error: ", err)
        }

        return result
    }

while loop?

    while continuePlaying == true
        prompt([]string)
        _, result, err := select{
            Label:
            Items:
        }

        if result == "quit" {
            
        }

menuSelections arr?

    var menuSelections []string
    
    raise
        timesCalled = 0
    call || check
        if timesCalled == 2 {
            menuSelections[1] = "check"
        }
    fold
    quit

#### gameplay

##### basic ai? :-/

node search algorithm?

fundamentally every option is a new node, the best "ai" would choose the most efficient route every time.

make lower difficulties by introducing random chance to choose a worse option?

##### win/loss

hierarchy of hands

eg:
    Royal Flush
    Flush
    Straight
    Full House
    etc...

higher cards - 
card value changes by game?

parse a json file of game values

    [
        {
            "name": "Texas Hold'em",
            "handSize": 5,
            "cardValues": {
                "Ace" 13,
		        "Two"1,
		        "Three"2,
                "Four"3,
                "Five"4,
                "Six"5,
                "Seven"6,
                "Eight"7,
                "Nine"8,
                "Ten"9,
                "Jack"10,
                "Queen"11,
                "King" 12,
            }
        }
    ]

    type game struct {
        name string
        handSize int
        cardValues []int
        // cardValues[0] = ace -> cardValues[12] = king
    }

    dealer := newDeck(cardValues)

refactor newDeck() to apply value to each card

    type deck []struct {
        card string
        rank int
    }

refactor saving and getting from file

refactor toString -> toJson?



#### player/dealer hands

define playerCount when starting game

    playerCount := menuPrompt()

function to append player to players slice?

    type player struct {
        name string
        hand deck
        chips int
    }

    players :=  []player

handSize defined by selecting game?

decks are the basis

    dealer := newDeck()

    dealer.shuffle()

    for i := range playerCount {
        players.append(
            player{
                name, 
                dealer.deal(handSize), 
                startChips
            }
        )
    }

#### betting

set stakes at start of game?

    var startChips int

    if stakes == "High Stakes" {
        startChips :=  2000
    }

## Stage two

i'm just kinda excitied to get some hands on experience with an actual game engine. i am
looking into learning C# at some point but its not my number one priority. SO ill get my
game dev itch scratched with GO.

### https://ebitengine.org/

ebitengine is apparently really simple which should be a fun first dip into actually working with
an engine.