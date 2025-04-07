# Here is the stuff I'm adding and understanding to the server

The most similar thing that I can describe an http server is an "event listener" in JavaScript
Waits for a request then delivers according to the request

The simple server is in the commit `1a417e0ce0b2c8096ff69f97dcb609788e1563c5`

#### Adding multiplexing

Is used to send multiple signals or streams of data over a single channel. It allows more efficient use of resources, such as bandwidth, by combining several signals into one. Once the combined signal reaches its destination, it is separated back into its original individual signals.

#### There are different types of multiplexing based on how the signals are combined: 

##### Time division multiplexing (TDM):
Each signal is assigned a specific time slot in the channel. Signals take turns using the channel
##### Frequency Division multiplexing (FDM):
Each signal is assigned a different Frequency within the available bandwidth
##### Wavelength Division multiplexing (WDM) (used in fiber optics):
Similar to FDM but uses different Wavelengths of light in optical fibers to carry multiple signals.
##### Code division multiplexing (CDM):
Different signals are assigned unique codes and transmitted simultaneously over the same Frequency.

#### Why multiplexing is important:
- Efficiency: It maximizes the utilization of available resources like bandwidth
- Cost-saving: Multiple signals can share a single physical infrastructure, reducing costs.
- Scalability: It allows for the system to grow by adding more signals or data streams without needing to add new physical channels.

#### Running multiple Servers at one time
Sometimes you may want to customize how the server runs, or you may want to run multiple HTTP servers in the same program at once. For example, you may have a public website and a private admin website you want to run from the same program. Since you can only have one default HTTP server, you wouldn't be able to do this with the default one.

