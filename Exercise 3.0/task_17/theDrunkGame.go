package thedrunkgame

import (
	"fmt"

	"github.com/kkeigenkai/Yandex/util"
)

func TheDrunkGame(fs, ss []int) string {
	fq, sq := util.CreateLinkedQueue(fs), util.CreateLinkedQueue(ss)
	cnt := 0

	for ; fq.Size != 0 && sq.Size != 0 && cnt < 1_000_000; cnt++ {
		if fq.Front() == 0 && sq.Front() == 9 {
			fq.Add(fq.Remove())
			fq.Add(sq.Remove())
		} else if sq.Front() == 0 && fq.Front() == 9 {
			sq.Add(fq.Remove())
			sq.Add(sq.Remove())
		} else if fq.Front() > sq.Front() {
			fq.Add(fq.Remove())
			fq.Add(sq.Remove())
		} else {
			sq.Add(fq.Remove())
			sq.Add(sq.Remove())
		}
	}

	res := fmt.Sprintf("first %d", cnt)
	if fq.Size == 0 {
		res = fmt.Sprintf("second %d", cnt)
	} else if cnt == 1_000_000 {
		res = "botva"
	}
	return res
}
