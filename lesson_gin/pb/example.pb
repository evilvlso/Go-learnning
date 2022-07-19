syntax="proto3";
option go_package=".;example";
package example;

message Teacher{
    string Name=1;
    int32 Age=2;
    enum Level {
       LECTURER=0;
       DOCTOR=1;
       MASTER=2;
    }
    Level level=3;
    string Course=4;
}
