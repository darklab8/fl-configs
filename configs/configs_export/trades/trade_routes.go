package trades

import "github.com/darklab8/fl-configs/configs/configs_mapped"

func CalculateTradeRoutes(configs *configs_mapped.MappedConfigs) {

	// for _, system := range configs.Systems.Systems {

	// }
}

// Algorithm should be like this:
// We iterate through list of Systems:
// Adding all bases, jump gates, jump holes, tradelanes as Vertexes.
// We scan in advance nicknames for object on another side of jump gate/hole and add it as vertix
// We calculcate distances between them. Distance between jump connections is 0 (or time to wait measured in distance)
// We calculate distances between trade lanes as shorter than real distance for obvious reasons.
// The matrix built on a fight run will be having connections between vertixes as hashmaps of possible edges? For optimized memory consumption in a sparse matrix.

// Then on second run, knowing amount of vertixes
// We build Floyd matrix? With allocating memory in bulk it should be rather rapid may be.
// And run Floud algorithm.
// Thus we have stuff calculated for distances between all possible trading locations. (edited)
// [6:02 PM]
// ====
// Then we build table of Bases as starting points.
// And on click we show proffits of delivery to some location. With time of delivery. And profit per time.
// [6:02 PM]
// ====
// Optionally print sum of two best routes that can be started within close range from each other.
