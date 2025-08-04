package generator

import "github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"

func GenerateRandomSampleFileName() string {
	return randomFilePrefix + "_" + utils.GetNowAsFormattedString() + ".txt"
}
