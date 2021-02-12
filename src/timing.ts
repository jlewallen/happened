export class Timer {
    constructor(private readonly id: number) {}

    static every(interval: number, callback: () => void): Timer {
        return new Timer(
            setInterval(async (): Promise<void> => {
                callback();
            }, interval)
        );
    }

    public cancel(): void {
        clearInterval(this.id);
    }
}
