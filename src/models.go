package main

import(
  "time"
)

type Sheet struct{
  id int;
  owner int;
  system int;
  table int;
  creationDate time.Time;
}
