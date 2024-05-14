package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"log"
)

func Union(cmd *cobra.Command, args []string) {
	srcHost, _ := cmd.Flags().GetString("src-host")
	srcPort, _ := cmd.Flags().GetInt("src-port")
	srcPassword, _ := cmd.Flags().GetString("src-password")
	dstHost, _ := cmd.Flags().GetString("dst-host")
	dstPort, _ := cmd.Flags().GetInt("dst-port")
	dstPassword, _ := cmd.Flags().GetString("dst-password")
	keyPattern, _ := cmd.Flags().GetString("key-pattern")
	log.Printf("src host is %s, src port is %d, src password is %s, dst host is %s, dst port is %d, dst password is %s\n",
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
		typeStr, err := srcRdb.Type(context.Background(), key).Result()
		if err != nil {
			log.Printf("get key %s type failed with error %v", key, err)
			continue
		}

		// copy if dst not exist
		exist, err := dstRdb.Exists(context.Background(), key).Result()
		if err != nil {
			log.Printf("asset key %s exist failed with error %v", key, err)
			continue
		}
		if exist > 0 {
			continue
		}

		log.Printf("copy key %s", key)
		// copy src to dst
		switch typeStr {
		case "zset":
			log.Printf("type is %s", typeStr)
			scores, err := srcRdb.ZRangeWithScores(context.Background(), key, 0, -1).Result()
			if err != nil {
				log.Printf("zrange 0 -1 failed with error %v", err)
				continue
			}
			for _, score := range scores {
				dstRdb.ZAdd(context.Background(), key, score)
			}
		case "hash":
			log.Printf("type is %s", typeStr)
			srcMap, err := srcRdb.HGetAll(context.Background(), key).Result()
			if err != nil {
				log.Printf("hgetall failed with error %v", err)
				continue
			}
			dstRdb.HMSet(context.Background(), key, srcMap)
		}
	}
}
