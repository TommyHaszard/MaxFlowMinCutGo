# MaxFlowMinCutGo

First project in Go!
This was just a little project turning my most recent uni assignment which i wrote in java, into Go so i could learn the language by recreating something fresh in my mind.

## How to run:
Super simple to run just use the make file by entering
### make
into the terminal.

## How to test different graphs:
I haven't implemented all the tests I did for the assignment yet, but its pretty simple if you want to make your own.
Create a new function and inside that init a new Graph by first creating the "Source" and "Sink" nodes.
Call "CreateVertex" method on these nodes to add them to the verticies map inside the Graph.
Call "AddEdge" method to create an edge between two nodes and specify the flow.
Then call the function in main and passed the Graph into either EdmondsKarp or FordFulkerson
