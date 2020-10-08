import {ITerminalAddon, Terminal} from 'xterm';
import {FitAddon} from 'xterm-addon-fit';
import 'xterm/css/xterm.css'

const bare = new Terminal({disableStdin: false});
const fitAddon = new FitAddon();
bare.loadAddon(<ITerminalAddon>fitAddon);


export class Xterm {
  elem: HTMLElement;
  term: Terminal;
  //decoder: lib.UTF8Decoder;

  message: HTMLElement;
  messageTimeout: number;


  constructor(elem: HTMLElement) {
    this.elem = elem;
    this.term = bare;
    // @ts-ignore
    this.message = elem.ownerDocument.createElement("div");
    this.message.className = "xterm-overlay";
    this.messageTimeout = 2000;

    this.term.open(elem);
    this.resize();
  };

  info(): { columns: number, rows: number } {
    return {columns: this.term.cols, rows: this.term.rows};
  };

  output(data: string) {
    this.term.write(data);
  };

  resize(): void {
    fitAddon.fit();
    this.term.scrollToBottom();
    console.log("resize")
  }


  removeMessage(): void {
    if (this.message.parentNode == this.elem) {
      this.elem.removeChild(this.message);
    }
  }

  setWindowTitle(title: string) {
    document.title = title;
  };

  setPreferences(value: object) {
  };

  onInput(callback: (input: string) => void) {
    this.term.onData((data) => {
      callback(data);
    });

  };

  onResize(callback: (colmuns: number, rows: number) => void) {
    this.term.onResize((data) => {
      callback(data.cols, data.rows);
    });
  };

  deactivate(): void {
    this.term.blur();
  }

  reset(): void {
    this.removeMessage();
    this.term.clear();
  }

  close(): void {
    this.term.dispose();
  }
}
