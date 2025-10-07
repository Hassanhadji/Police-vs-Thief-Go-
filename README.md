The board size is randomized: rows ∈ [10, 200], cols ∈ [10, 200].

Two players:

Police — starts at (0,0) with a random move budget.

Thief — starts at the bottom-right (rows-1, cols-1).

Each “turn,” the controller:

Calls thiefplays()

Calls policeplays()

Sends status messages via channels to each player

Prints current positions

The game ends when:

Police catches the thief (same cell) → Police wins

Police runs out of moves without catching → Thief wins

Both end up at (0,0) simultaneously → Tie

Thief reaches (0,0) → Thief escapes
