package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/gin-gonic/gin"
  "log"
  "os"
  "time"
)

var (
  Conexion = "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASS") + "@" + os.Getenv("POSTGRES_HOST") + "/asterisk?sslmode=disable"
  AsteriskDB *sql.DB
)

type registro struct {
  Id int `json:"id"`
  Start time.Time `json:"timestamp"`
  Src string `json:"origen"`
  Dst string `json:"destino"`
  Dstchannel string `json:"atendio"`
  Duration int `json:"duracion"`
  Disposition string `json:"estado"`
  Uniqueid string `json:"uniqueid"`
}

func init() {
  var err error
  AsteriskDB, err = sql.Open("postgres", Conexion)
  if err != nil {
    log.Fatal(err)
  }
}

func barraLlamadas(c *gin.Context) {
  var err error
  var rs []*registro

  rows, err := AsteriskDB.Query("SELECT id,start,src,dst,dstchannel,duration,disposition,uniqueid FROM cdr WHERE disposition='ANSWERED'")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    r := new(registro)

    err = rows.Scan(&r.Id, &r.Start, &r.Src, &r.Dst, &r.Dstchannel, &r.Duration, &r.Disposition, &r.Uniqueid)
    if err != nil {
      log.Fatal(err)
    }

    rs = append(rs, r)
  }

  c.JSON(200, rs)
}

func barraLlamadasDesde(c *gin.Context) {
  var err error
  var rs []*registro

  id := c.Param("id")

  rows, err := AsteriskDB.Query("SELECT id,start,src,dst,dstchannel,duration,disposition,uniqueid FROM cdr WHERE disposition='ANSWERED' AND src='" + id + "'")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    r := new(registro)

    err = rows.Scan(&r.Id, &r.Start, &r.Src, &r.Dst, &r.Dstchannel, &r.Duration, &r.Disposition, &r.Uniqueid)
    if err != nil {
      log.Fatal(err)
    }

    rs = append(rs, r)
  }

  c.JSON(200, rs)
}

func barraLlamadasId(c *gin.Context) {
  var err error
  var rs []*registro

  id := c.Param("id")

  rows, err := AsteriskDB.Query("SELECT id,start,src,dst,dstchannel,duration,disposition,uniqueid FROM cdr WHERE disposition='ANSWERED' AND uniqueid='" + id + "'")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    r := new(registro)

    err = rows.Scan(&r.Id, &r.Start, &r.Src, &r.Dst, &r.Dstchannel, &r.Duration, &r.Disposition, &r.Uniqueid)
    if err != nil {
      log.Fatal(err)
    }

    rs = append(rs, r)
  }

  c.JSON(200, rs)
}

func barraLlamadasFecha(c *gin.Context) {
  var err error
  var rs []*registro

  id := c.Param("id")

  rows, err := AsteriskDB.Query("SELECT id,start,src,dst,dstchannel,duration,disposition,uniqueid FROM cdr WHERE disposition='ANSWERED' AND start::text LIKE '" + id + "%'")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    r := new(registro)

    err = rows.Scan(&r.Id, &r.Start, &r.Src, &r.Dst, &r.Dstchannel, &r.Duration, &r.Disposition, &r.Uniqueid)
    if err != nil {
      log.Fatal(err)
    }

    rs = append(rs, r)
  }

  c.JSON(200, rs)
}


func main() {
  r := gin.Default()
  r.GET("/llamadas", barraLlamadas)
  //r.GET("/llamadas/id/:id", barraLlamadasId)
  //r.GET("/llamadas/origen/:id", barraLlamadasDesde)
  //r.GET("/llamadas/fecha/:id", barraLlamadasFecha)
  r.Run()
}
