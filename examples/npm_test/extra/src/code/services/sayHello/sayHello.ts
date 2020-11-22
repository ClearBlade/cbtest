
interface Request {
    params: Record<string, string>;
}

interface Response {
    success(message: string);
}

function makeHelloMessage(name: string): string {
    return `Hello, ${name}`;
}

function sayHello(req: Request, resp: Response) {

    const { name = "world!" } = req.params;

    resp.success(makeHelloMessage(name));
}
