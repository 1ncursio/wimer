class Timer {
  private timer: NodeJS.Timeout | null = null;

  constructor(private callback: () => void, private tick: number) {}

  start() {
    if (this.timer) {
      return;
    }
    this.timer = setInterval(this.callback, this.tick);
  }

  stop() {
    if (!this.timer) {
      return;
    }
    clearInterval(this.timer);
    this.timer = null;
  }

  reset() {
    this.stop();
  }
}

export default Timer;
