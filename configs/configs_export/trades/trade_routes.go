package trades

import (
	"math"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/conftypes"
)

type SystemObject struct {
	nickname string
	pos      conftypes.Vector
}

func DistanceForVecs(Pos1 conftypes.Vector, Pos2 conftypes.Vector) float64 {
	// if _, ok := Pos1.X.GetValue(); !ok {
	// 	return 0, errors.New("no x")
	// }
	// if _, ok := Pos2.X.GetValue(); !ok {
	// 	return 0, errors.New("no x")
	// }

	x_dist := math.Pow((Pos1.X - Pos2.X), 2)
	y_dist := math.Pow((Pos1.Y - Pos2.Y), 2)
	z_dist := math.Pow((Pos1.Z - Pos2.Z), 2)
	distance := math.Pow((x_dist + y_dist + z_dist), 0.5)
	return distance
}

func MapConfigsToFloyder(configs *configs_mapped.MappedConfigs) *Floyder {
	floyder := NewFloyder()
	for _, system := range configs.Systems.Systems {

		var system_objects []SystemObject = make([]SystemObject, 0, 50)

		for _, system_obj := range system.Bases {
			object := SystemObject{
				nickname: system_obj.Base.Get(),
				pos:      system_obj.Pos.Get(),
			}

			for _, existing_object := range system_objects {
				distance := DistanceForVecs(object.pos, existing_object.pos)
				floyder.SetEdge(object.nickname, existing_object.nickname, distance)
			}

			system_objects = append(system_objects, object)
		}

		for _, jumphole := range system.Jumpholes {
			object := SystemObject{
				nickname: jumphole.Nickname.Get(),
				pos:      jumphole.Pos.Get(),
			}

			for _, existing_object := range system_objects {
				distance := DistanceForVecs(object.pos, existing_object.pos)
				floyder.SetEdge(object.nickname, existing_object.nickname, distance)
			}

			jumphole_target_hole := jumphole.GotoHole.Get()
			floyder.SetEdge(object.nickname, jumphole_target_hole, 0)
			system_objects = append(system_objects, object)
		}
	}

	return floyder
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
