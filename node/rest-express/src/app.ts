import * as bodyParser from 'body-parser'
import express from 'express'

const characters =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
const charactersLength = characters.length

const app = express()

var jsonParser = bodyParser.json({
    limit: '50mb',
    verify(req: any, res, buf, encoding) {
        req.rawBody = buf
    }
})

app.post('/', jsonParser, (req: express.Request, res: express.Response) => {
    const len = Math.floor(Math.random() * 800 + 200)
    const out = [...Array(len)].map((_) =>
        String.fromCharCode(characters.charCodeAt(
            Math.floor(Math.random() * charactersLength))))

    res.send({ result: req.body.name + out.join('') })
})

export { app }
