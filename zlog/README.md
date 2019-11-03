# Lab 7: Processing Channel Zaps

| Lab 7:                | Processing Channel Zaps                          |
| --------------------- | ------------------------------------------------ |
| Subject:              | DAT320 Operating Systems and Systems Programming |
| Deadline:             | **November 19, 2019**                            |
| Expected effort:      | 15-20 hours                                      |
| Grading:              | Graded                                           |
| Submission:           | Group                                            |

## Table of Contents

1. [Introduction](#introduction)
2. [Collecting Channel Zaps](#collecting-channel-zaps)
3. [Traffic Generator](#traffic-generator)
4. [Building a Zap Event Processing Server](#building-a-zap-event-processing-server)
5. [Publish/Subscribe RPC Client and Server](#publishsubscribe-grpc-client-and-server)

## Introduction

In this larger project, you will work in groups of 2 or 3 students.

**Heads up!** Before you begin, you should read through the whole document, as this will help you plan your design.
Specifically, you may want to think through how you separate code into different files, name your packages and structs, and so forth to design a good piece of software.

**Another Heads up!** You are expected to use the provided template code.
Replace `TODO` comments when you have implemented your solution with `DONE`.
When submitting, make sure that your program compiles and runs.
If have trouble with some of tasks below:
(1) Try to get help during the lab exercise hours;
(2) If you are unable to get help in time for the deadline, simply comment the non-compiling code, and provide a commit comment describing the problem.
This comment should be separate from the first line of the final commit message, which should say:

`username lab7 submission`

## Collecting Channel Zaps

Imagine that you are working for ZapBox, an Internet and Cable service provider.
ZapBox has deployed a large number of set-top boxes at customer homes that allows them to watch TV over a fiber optic cable.
The TV signal is distributed to customers based on a multicast stream for each available TV channel in ZapBox’s channel portfolio.

Recently, ZapBox commissioned a software update on their deployed set-top boxes.
After this software update, the set-top box will send a UDP message to a server every time a user changes the channel on their set-top box.
In addition to channel changes, a few other items of interest may also be sent.
Thus, a message sent by a set-top box may contain information about either channel changes, volume, mute status, or HDMI status.
The content depends on the actions of the different TV viewers.
Below is shown a few samples of the message format:

```log
2013/07/20, 21:56:13, 252.126.91.56, HDMI_Status: 0
2013/07/20, 21:56:55, 111.229.208.129, MAX, Viasat 4
2013/07/20, 21:57:48, 98.202.244.97, FEM, TVNORGE
2013/07/20, 21:57:44, 12.23.36.158, Canal 9, MAX
2013/07/20, 21:57:46, 81.187.186.219, TV2 Bliss, TV2 Zebra
2013/07/20, 21:57:42, 61.77.4.101, TV2 Film, TV2 Bliss
2013/07/20, 21:57:42, 203.124.29.72, Volume: 50
2013/07/20, 21:57:42, 203.124.29.72, Mute_Status: 0
```

Each line above represents an event, triggered by a single TV viewer's action, either to change the channel on their set-top box, or adjust the volume and so forth.
These set-top box events are sent in the text format shown above.
The fields are separated by comma and have the meaning shown in the table below.
A message format with 5 fields represents channel change events.
A message with only 4 fields contains a status change in the 4th field, and no 5th field.

| Field No. | Field Name   | Description                                       |
| --------- | ------------ | ------------------------------------------------- |
| 1         | Date         | The date that the event was sent.                 |
| 2         | Time         | The time that the event was sent.                 |
| 3         | IP           | The IPv4 address of the sending set-top box unit. |
| 4         | FromChan     | The previous channel of the set-top box.          |
| 5         | ToChan       | The new channel of the set-top box.               |
| 4         | StatusChange | A change in status on the set-top box.            |

A `StatusChange` may contain one of the following entries:

| StatusChange | Value range | Description                                                                                         |
| ------------ | ----------- | --------------------------------------------------------------------------------------------------- |
| Volume:      | 0-100       | The volume setting on the set-top box.                                                              |
| Mute Status: | 0/1         | The mute setting on the set-top box.                                                                |
| HDMI Status: | 0/1         | The HDMI status of the set-top box indicates whether or not it is powered on and connected to a TV. |

## Traffic Generator

For the purposes of this lab project, we have built a traffic generator to simulate the events generated by the set-top boxes.
The traffic generator resends set-top box events loaded from a large dataset obtained from real traffic.
The IP addresses have been scrambled and do not represent a real set-top box.
The traffic generator works by synchronizing the timestamp obtained from the dataset with the local clock on the simulator machine.
The date is not synchronized.

In a real deployment, the traffic would typically be sent from set-top boxes using UDP and received at a single UDP server, where the data can be processed.
However, to make the simulator scale to multiple receiver groups (you the students), we have instead set up the traffic generator on a single machine multicasting each set-top box event to a single multicast address.

## Part 1: Building a Zap Event Processing Server

The objective of this part is to develop a UDP multicast server that will process the events that are sent by the set-top boxes (in our case the traffic generator).
Your UDP server can run on one of the machines in the Linux lab.
Your server should be able to receive UDP packets from the traffic generator using multicast address and port:

`224.0.1.130:10000`

Note that since the traffic generator is continuously sending out a stream of zap events, it may be difficult to work with this part of the lab on your own machine.
The multicast stream is only available on the subnet of the Linux lab.
It is therefore recommended that you work on the lab machines, either physically in the Linux lab (E353), or remotely using ssh.

## Tasks Part 1 (60 points)

1. (5 points) Build a UDP zapserver that listens to the IP multicast address and port number specified above.
   Your server *must not* echo anything back (respond) to the traffic generator.
   Your server should only receive zap events in a loop.
   In this task, the server only needs to print to the console whatever it receives from the traffic generator.
   Hint: `net.ListenMulticastUDP`.

   To run your code on the Linux lab with ssh we suggest the following workflow:

   ```console
   go build
   ```

   Or if you are working on Windows (with WSL), you may need to cross-compile the code as follows:

   ```console
   GOOS=linux GOARCH=amd64 go build
   ```

   Finally, copy the binary to a unix machine:

   ```console
   scp zapserver <username>@ssh1.ux.uis.no:bin/
   ```

   Replace `<username>` with your username on the Unix system.
   Note that you may need to create the `bin` folder in your `$HOME` directory before running scp.
   Also, you can avoid the `username` by creating a `.ssh/config` file:

   ```console
   Host *.ux.uis.no
     User <username>
     IdentityFile ~/.ssh/id_rsa
   ```

   Further, you may also wish to add `bin` to your `$PATH`, whereever it is specified.

   When you have copied the `zapserver` to the Unix system, you can run login to another server, e.g. pi25, and run it.

   ```console
   ssh pi25.ux.uis.no
   zapserver
   ```

2. (10 points) Develop a data structure for storing individual zap events.
   The struct must contain all the necessary fields to store channel changes (ignore storing status changes for now).
   The test cases in `chzap_test.go` should pass.
   The main task here is to implement the constructor `NewSTBEvent()` which can be used by your server when it receives a zap event.
   In addition the struct should have the following methods associated methods.
   See the template in `chzap.go`.

   | Method                                 | Description                                                                 |
   | -------------------------------------- | --------------------------------------------------------------------------- |
   | NewSTBEvent()                          | Returns either a channel zap event, a status change event, or an error.     |
   | String() string                        | Return a string representation of your struct.                              |
   | Duration(provided ChZap) time.Duration | Return the duration between the receiving zap event and the provided event. |

   Hints: `time.Time package`, Methods: `time.Parse()`, `strings.Split()`, `strings.TrimSpace()`,
   Layout: `const timeLayout = "2006/01/02, 15:04:05"`

3. (10 points) The next task is to use `zlog/simplelogger.go` to store the channel changes received on your zapserver.
   a) Use the API of the simple logger to compute the number of viewers on `NRK1` periodically, once every second.
   Print the output to the console.
   b) Implement the same for `TV2 Norge`.
   They should both be printed to the console on a separate line.

4. (5 points) Measure the time it takes to compute the `Viewers()` function using `TimeElapsed()`.
   Take note of the measurements obtained for the `Viewers()` function over time.
   You may also monitor the memory usage at runtime using `runtime.ReadMemStat()`.
   What does these results show?
   What could be the cause of the observed problem?

5. (10 points) Implement a function that can compute a list of the top-10 channels.
   Call this function periodically, once every second.
   Hint: Results returned from the `ChannelViewers()` method defined in the `ZapLogger` interface can be sorted.

   *Note that the underlying data structure used so far precludes an efficient implementation.*

6. (10 points) Implement a new data structure `ViewersLog` that avoids the problems that you should have identified with the simple slice-based storage solution.
   Implement the data structure so that it can support you with keeping track of the top-10 list of channels.
   Your implementation must adhere to the `ZapLogger` interface.
   Hint: You do not need to store all the zap events to compute the number of viewers for each channel.

7. (10 points) To ensure that the viewer count for each channel is completely accurate, implement another data structure `ConcurrentLog` that facilitates fine-grained locking per TV channel.

## Part 2: Publish/Subscribe gRPC Client and Server (40 points)

In this part you will implement a gRPC-based server that accepts subscribe requests from external clients.
After receiving a subscribe request, the server is expected to send a stream of viewership statistics (the top-10 list) to the subscribed clients.
The gRPC server should implement the `Subscribe()` method defined in the `subscribe.proto` file.

A client invoking the subscribe request should at least include the following fields in the subscribe request.

* The `ClientAddr` specifies the IP address and port number of a subscribed client.
* The `RefreshRate` specifies how often a subscribed client wishes to be notified.

Before you begin, you will want to install a few dependencies.

To install the protobuf compiler, `protoc`, navigate [here](https://github.com/google/protobuf/releases) and download the most recent version.
Next, you also need to get the protobuf compiler plugin for Go:

//TODO check if these are downloaded by automatically when/if go.mod file specifies these

```console
go get github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/golang/protobuf/{proto,protoc-gen-go} // TODO check: try this if the above doesn't work??
go get google.golang.org/grpc
```

1. (10 points) Implement the gRPC server in a separate file (`publisher.go`).
   However, it should be compiled together with the UDP zapserver from Part 1.

   * The gRPC server should serve statistics based on zap events from the zap storage.
   * This should be done while continuously updating the server's storage (the state of the server).
   * The gRPC server and the zapserver part receiving zap events should be implemented as separate goroutines.
   * Assume that the refresh rate is one second or more.

   The included `subscribe.proto` file is provided for you.
   You will need to add the necessary variables to the messages.
   Note that the `stream` keyword specified in the response message for the `Subscribe()` method.
   Streaming is needed here, since we want to update the client periodically without the client pulling from the server.

   Once you have finished defining the content of the `subscribe.proto` file, you must generate the corresponding Go code as follows:

   ```console
   protoc --go_out=plugins=grpc:. subscribe.proto
   ```

2. (10 points) Next implement a command line subscriber gRPC client in a separate file (`subscriber.go`).
   The client should be compiled into a separate binary.
   The client should display the viewership updates as they are received from the gRPC server.
   Note that it is the server's responsibility to periodically send updates to the subscribed clients.

3. (2 points) How would you characterize the access pattern hitting the server's state?
   That is, what is the relationship between reads and writes to the server's state (storage datastructure).
   You should make a drawing of the system architecture, illustrating the relationship between the gRPC clients and the server (i.e. the combined gRPC server and zapserver) and the STBs and the server.

4. (3 points) With this access pattern (workload) in mind.
   How would you protect the server's state to avoid returning a statistics computation that is incorrect or otherwise malformed?

5. (10 points) Now we want to analyze the duration between channel change clicks.
   To do that, we need to store the previous zap event for each IP, so that you can use the `Duration()` method that you developed earlier.
   You will need to extend your new data structure or add another data structure for storing these durations.
   Also, extend the `SubscribeRequest` message with an additional field to select the type of statistics the subscription refers to.
   Two types should be supported: viewership and duration statistics.
   Whatever statistics is chosen by the subscriber, your publisher should send publications to the relevant subscribers at the specified refresh rate.

6. (5 points) Profile the data structure implemented in Part 2.1.
   Implement a data structure that better supports the workload experienced by the zapserver.
   Profile the new data structure and compare the results to the one implemented in Part 2.1.
