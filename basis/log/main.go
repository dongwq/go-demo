/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-8，下午1:17
 * @version 0.1
 */
package main


import log "github.com/cihub/seelog"

func main() {
	defer log.Flush()
	log.Info("Hello from Seelog!")
}
