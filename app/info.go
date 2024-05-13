package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"log"
)

func Info(cmd *cobra.Command, args []string) {
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

	printRDB(srcRdb, "src", keyPattern)
	printRDB(dstRdb, "dst", keyPattern)
}

func printRDB(rdb *redis.Client, tag string, keyPattern string) {
	infoKey := "redis-tools-info"
	keys, _ := rdb.Keys(context.Background(), keyPattern).Result()
	log.Printf("keys count is %d", len(keys))
	for _, key := range keys {
		str, _ := rdb.Type(context.Background(), key).Result()
		rdb.ZIncrBy(context.Background(), infoKey, 1, str)
	}

	info, _ := rdb.ZRangeWithScores(context.Background(), infoKey, 0, -1).Result()
	log.Printf("rdb<%s> info is %v", tag, info)

	rdb.Del(context.Background(), "redis-tools-info")
}
