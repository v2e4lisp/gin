package gin_test

import (
        "testing"

        gin "."
)

func Test_Watch(t *testing.T) {
        err := gin.LoadWatchPatterns("test_fixtures/.ginwatch")

        expect(t, err, nil)
        // Default. Watch go  file
        expect(t, gin.IsWatched("gofile.go"), true)
        // Default. Do not watch hidden files
        expect(t, gin.IsWatched(".hiddenfile"), false)
        // Patterns specified in .ginwatch file
        expect(t, gin.IsWatched("watch/test.tmpl"), true)
        // Patterns not specified in .ginwatch file
        expect(t, gin.IsWatched("ignore/test.tmpl"), false)
}
