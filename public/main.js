// overlay
const time = document.getElementById('time');
const word = document.getElementById('word');
const scr = document.getElementById('screen');

let screenTimeout = null;
async function showScreen(type) {
    if (screenTimeout) {
        clearTimeout(screenTimeout);
        resetScreen();
        await new Promise((r) => setTimeout(r, 200));
    }

    /** @type {HTMLAudioElement} */
    const audio = document.getElementById('audio-' + type);
    scr.style.background = type === 'wrong' ? '#f44336' : '#32cd32';
    scr.style.animationIterationCount = type === 'wrong' ? 4 : 3;
    scr.classList.add('screen');

    screenTimeout = setTimeout(resetScreen, scr.style.animationIterationCount * 1e3);

    audio.pause();
    audio.currentTime = 0;
    audio.play();
}

function resetScreen() {
    scr.style.background = '';
    scr.style.animationIterationCount = 0;
    scr.classList.remove('screen');
    playingScreen = null;
}

// ws
const url = new URL(window.location);
url.pathname = '/ws';
url.protocol = url.protocol === 'https:' ? 'wss:' : 'ws:';

let connectInterval;
function connect() {
    var ws = new WebSocket(url);
    window.ws = ws;
    ws.onopen = () => {
        if (connectInterval) {
            clearInterval(connectInterval);
            connectInterval = null;
        }
        console.log('[WS]: Connected');
    };

    ws.onmessage = ({ data }) => {
        try {
            const parsed = JSON.parse(data);
            if ('time' in parsed)
                time.innerText = parseTime(parsed.time);
            if ('screen' in parsed)
                showScreen(parsed.screen);
            if ('word' in parsed)
                word.innerText = parsed.word;
        } catch (e) {
            console.error('[WS]: failed to parse message', data, e);
        }
    };

    ws.onclose = () => {
        if (!connectInterval)
            connectInterval = setInterval(() => {
                console.log('[WS]: Trying to reconnect...');
                connect();
            }, 5e3);
        console.log('[WS]: Disconnected');
    };
}

connect();

function parseTime(seconds) {
    if (typeof seconds !== 'number')
        return '00:00';

    return `${Math.floor(seconds / 60).toString().padStart(2, '0')}:${(seconds % 60).toString().padStart(2, '0')}`;
}