import { RemoteGraphQLDataSource } from '@apollo/gateway'

export class ModifyRequestAndResponse extends RemoteGraphQLDataSource {
  willSendRequest({ request, context }) {
    // Pass the user's id from the context to underlying services
    // as a header called `user-id`
    request.http.headers.set('x-access-token', context.userId)
  }

  async didReceiveResponse({ response, request, context }) {
    // Parse the Server-Id header and add it to the array on context
    const serverId = response.http.headers.get('Server-Id')
    if (serverId) {
      context.serverIds.push(serverId)
    }
    // Return the response, even when unchanged.
    return response
  }
}
