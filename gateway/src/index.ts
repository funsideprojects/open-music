import { ApolloServer } from 'apollo-server'
import { ApolloGateway } from '@apollo/gateway'

import { ModifyRequestAndResponse } from './middlewares'

const gateway = new ApolloGateway({
  serviceList: [{ name: 'identity', url: 'http://localhost:4001/query' }],
  buildService({ name, url }) {
    console.log('buildService -> name', name)

    return new ModifyRequestAndResponse({ url })
  },
})

const server = new ApolloServer({
  gateway,
  // Subscriptions are not currently supported in Apollo Federation
  subscriptions: false,
  context: ({ req }) => {
    let userId: string = undefined
    const accessToken = req.headers['x-access-token']

    if (accessToken) {
      userId = 'abcd'
    }

    // Add the user ID to the context
    return { req, userId }
  },
})

server
  .listen()
  .then(({ url }) => {
    console.log(`Server ready at ${url}`)
  })
  .catch((err) => {
    console.error(err)
  })
