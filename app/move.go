package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"log"
)

func Move(cmd *cobra.Command, args []string) {
	srcHost, _ := cmd.Flags().GetString("src-host")
	srcPort, _ := cmd.Flags().GetInt("src-port")
	srcPassword, _ := cmd.Flags().GetString("src-password")
	dstHost, _ := cmd.Flags().GetString("dst-host")
	dstPort, _ := cmd.Flags().GetInt("dst-port")
	dstPassword, _ := cmd.Flags().GetString("dst-password")
	keyPattern, _ := cmd.Flags().GetString("key-pattern")
	fmt.Printf("src host is %s, src port is %d, src password is %s, dst host is %s, dst port is %d, dst password is %s\n",
		srcHost, srcPort, srcPassword, dstHost, dstPort, dstPassword)

	srcRdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", srcHost, srcPort),
		Password: srcPassword, // no password set
		DB:       0,           // use default DB
	})

	dstRdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", dstHost, dstPort),
		Password: dstPassword, // no password set
		DB:       0,           // use default DB
	})

	keys, _ := srcRdb.Keys(context.Background(), keyPattern).Result()

	for _, key := range keys {
		scores, _ := srcRdb.ZRangeWithScores(context.Background(), key, 0, -1).Result()
		for _, score := range scores {
			val := score.Score
			mem := score.Member.(string)
			log.Printf("inc src val is %f, mem is %s", val, mem)
			dstRdb.ZIncrBy(context.Background(), key, val, mem)
			srcRdb.ZRem(context.Background(), key, mem)
		}
	}
}
