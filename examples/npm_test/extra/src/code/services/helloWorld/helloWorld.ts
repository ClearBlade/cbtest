
interface Request {
    params: Record<string, string>;
}

interface Response {
    success(message: string);
}

function sayHello(name: string): string {
    return `Hello, ${name}`;
}

function helloWorld(req: Request, resp: Response) {

    const { name = "world!" } = req.params;

    resp.success(sayHello(name));
}
