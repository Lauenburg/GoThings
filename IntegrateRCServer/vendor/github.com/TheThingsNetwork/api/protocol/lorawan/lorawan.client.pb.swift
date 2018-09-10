/*
 * DO NOT EDIT.
 *
 * Generated by the protocol buffer compiler.
 * Source: github.com/TheThingsNetwork/api/protocol/lorawan/device_address.proto
 *
 */

/*
 * Copyright 2017, gRPC Authors All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import Foundation
import Dispatch
import gRPC

/// Type for errors thrown from generated client code.
internal enum Lorawan_DevAddrManagerClientError : Error {
  case endOfStream
  case invalidMessageReceived
  case error(c: CallResult)
}

/// GetPrefixes (Unary)
internal class Lorawan_DevAddrManagerGetPrefixesCall {
  private var call : Call

  /// Create a call.
  fileprivate init(_ channel: Channel) {
    self.call = channel.makeCall("/lorawan.DevAddrManager/GetPrefixes")
  }

  /// Run the call. Blocks until the reply is received.
  fileprivate func run(request: Lorawan_PrefixesRequest,
                       metadata: Metadata) throws -> Lorawan_PrefixesResponse {
    let sem = DispatchSemaphore(value: 0)
    var returnCallResult : CallResult!
    var returnResponse : Lorawan_PrefixesResponse?
    _ = try start(request:request, metadata:metadata) {response, callResult in
      returnResponse = response
      returnCallResult = callResult
      sem.signal()
    }
    _ = sem.wait(timeout: DispatchTime.distantFuture)
    if let returnResponse = returnResponse {
      return returnResponse
    } else {
      throw Lorawan_DevAddrManagerClientError.error(c: returnCallResult)
    }
  }

  /// Start the call. Nonblocking.
  fileprivate func start(request: Lorawan_PrefixesRequest,
                         metadata: Metadata,
                         completion: @escaping (Lorawan_PrefixesResponse?, CallResult)->())
    throws -> Lorawan_DevAddrManagerGetPrefixesCall {

      let requestData = try request.serializedData()
      try call.start(.unary,
                     metadata:metadata,
                     message:requestData)
      {(callResult) in
        if let responseData = callResult.resultData,
          let response = try? Lorawan_PrefixesResponse(serializedData:responseData) {
          completion(response, callResult)
        } else {
          completion(nil, callResult)
        }
      }
      return self
  }
}

/// GetDevAddr (Unary)
internal class Lorawan_DevAddrManagerGetDevAddrCall {
  private var call : Call

  /// Create a call.
  fileprivate init(_ channel: Channel) {
    self.call = channel.makeCall("/lorawan.DevAddrManager/GetDevAddr")
  }

  /// Run the call. Blocks until the reply is received.
  fileprivate func run(request: Lorawan_DevAddrRequest,
                       metadata: Metadata) throws -> Lorawan_DevAddrResponse {
    let sem = DispatchSemaphore(value: 0)
    var returnCallResult : CallResult!
    var returnResponse : Lorawan_DevAddrResponse?
    _ = try start(request:request, metadata:metadata) {response, callResult in
      returnResponse = response
      returnCallResult = callResult
      sem.signal()
    }
    _ = sem.wait(timeout: DispatchTime.distantFuture)
    if let returnResponse = returnResponse {
      return returnResponse
    } else {
      throw Lorawan_DevAddrManagerClientError.error(c: returnCallResult)
    }
  }

  /// Start the call. Nonblocking.
  fileprivate func start(request: Lorawan_DevAddrRequest,
                         metadata: Metadata,
                         completion: @escaping (Lorawan_DevAddrResponse?, CallResult)->())
    throws -> Lorawan_DevAddrManagerGetDevAddrCall {

      let requestData = try request.serializedData()
      try call.start(.unary,
                     metadata:metadata,
                     message:requestData)
      {(callResult) in
        if let responseData = callResult.resultData,
          let response = try? Lorawan_DevAddrResponse(serializedData:responseData) {
          completion(response, callResult)
        } else {
          completion(nil, callResult)
        }
      }
      return self
  }
}

/// Call methods of this class to make API calls.
internal class Lorawan_DevAddrManagerService {
  private var channel: Channel

  /// This metadata will be sent with all requests.
  internal var metadata : Metadata

  /// This property allows the service host name to be overridden.
  /// For example, it can be used to make calls to "localhost:8080"
  /// appear to be to "example.com".
  internal var host : String {
    get {
      return self.channel.host
    }
    set {
      self.channel.host = newValue
    }
  }

  /// Create a client that makes insecure connections.
  internal init(address: String) {
    gRPC.initialize()
    channel = Channel(address:address)
    metadata = Metadata()
  }

  /// Create a client that makes secure connections.
  internal init(address: String, certificates: String?, host: String?) {
    gRPC.initialize()
    channel = Channel(address:address, certificates:certificates, host:host)
    metadata = Metadata()
  }

  /// Synchronous. Unary.
  internal func getprefixes(_ request: Lorawan_PrefixesRequest)
    throws
    -> Lorawan_PrefixesResponse {
      return try Lorawan_DevAddrManagerGetPrefixesCall(channel).run(request:request, metadata:metadata)
  }
  /// Asynchronous. Unary.
  internal func getprefixes(_ request: Lorawan_PrefixesRequest,
                  completion: @escaping (Lorawan_PrefixesResponse?, CallResult)->())
    throws
    -> Lorawan_DevAddrManagerGetPrefixesCall {
      return try Lorawan_DevAddrManagerGetPrefixesCall(channel).start(request:request,
                                                 metadata:metadata,
                                                 completion:completion)
  }
  /// Synchronous. Unary.
  internal func getdevaddr(_ request: Lorawan_DevAddrRequest)
    throws
    -> Lorawan_DevAddrResponse {
      return try Lorawan_DevAddrManagerGetDevAddrCall(channel).run(request:request, metadata:metadata)
  }
  /// Asynchronous. Unary.
  internal func getdevaddr(_ request: Lorawan_DevAddrRequest,
                  completion: @escaping (Lorawan_DevAddrResponse?, CallResult)->())
    throws
    -> Lorawan_DevAddrManagerGetDevAddrCall {
      return try Lorawan_DevAddrManagerGetDevAddrCall(channel).start(request:request,
                                                 metadata:metadata,
                                                 completion:completion)
  }
}
