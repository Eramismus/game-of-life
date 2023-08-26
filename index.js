let game_data = null
let ws = null
function sendStuff() {
    if (ws.readyState == WebSocket.OPEN) {
        console.log('Sending message')
        ws.send('move_forward')
    }
}

function initialize() {
    if ('WebSocket' in window) {
        ws = new WebSocket('ws://localhost:10000/endpoint')
        ws.onopen = function () {
            console.log('Connected to server')
        }
        ws.onmessage = function (evt) {
            const received_msg = evt.data
            game_data = JSON.parse(received_msg)
            createTable()
        }
        ws.onclose = function () {
            // websocket is closed.
            alert('Connection is closed...')
        }
    } else {
        // The browser doesn't support WebSocket
        alert('WebSocket NOT supported by your Browser!')
    }
}

function createTable() {
    gridContainer = document.getElementById('gridContainer')
    if (!gridContainer) {
        // Throw error
        console.error('Problem: No div for the drid table!')
    }
    if (!document.getElementById('grid')) {
        table = document.createElement('table')
        table.setAttribute('id', 'grid')

        for (i = 0; i < game_data.grid.length; i++) {
            tr = document.createElement('tr')
            for (j = 0; j < game_data.grid[i].length; j++) {
                cell = document.createElement('td')
                cell.setAttribute('id', i + '_' + j)
                if (game_data.grid[i][j] == true) {
                    cell.setAttribute('class', 'live')
                } else {
                    cell.setAttribute('class', 'dead')
                }
                tr.appendChild(cell)
            }
            table.appendChild(tr)
        }
        gridContainer.appendChild(table)
    } else {
        for (i = 0; i < game_data.grid.length; i++) {
            for (j = 0; j < game_data.grid[i].length; j++) {
                cell = document.getElementById(i + '_' + j)
                if (game_data.grid[i][j] == false) {
                    cell.setAttribute('class', 'dead')
                } else {
                    cell.setAttribute('class', 'live')
                }
            }
        }
    }
}
// Start everything
window.onload = initialize()
