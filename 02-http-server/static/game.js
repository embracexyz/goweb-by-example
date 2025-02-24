let secretNumber = Math.floor(Math.random() * 100) + 1;
let attempts = 0;

function checkGuess() {
    const guess = parseInt(document.getElementById('guess').value);
    attempts++;

    if (isNaN(guess) || guess < 1 || guess > 100) {
        document.getElementById('result').innerText = "请输入一个 1 到 100 之间的数字";
        return;
    }

    if (guess === secretNumber) {
        document.getElementById('result').innerText = `恭喜你！猜对了！你用了 ${attempts} 次机会。`;
    } else if (guess < secretNumber) {
        document.getElementById('result').innerText = "猜的数字太小了，再试一次！";
    } else {
        document.getElementById('result').innerText = "猜的数字太大了，再试一次！";
    }
}

