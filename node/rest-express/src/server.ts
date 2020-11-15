import { AddressInfo } from 'net'
import { app } from './app.js'

const server = app.listen(50051, '0.0.0.0', () => {
    const { port, address } = server.address() as AddressInfo
    console.log('Server listening on:', 'http://' + address + ':' + port)
})
