(async () => {
  function findOrCreateBlock() {
    const rightBlock = document.getElementById('rhs');
    if (rightBlock) {
      return rightBlock;
    }

    const searchPage = document.getElementById('rcnt');
    const block = document.createElement('div');
    block.id = 'rhs';
    block.role = 'complementary';
    searchPage.append(block);
    return block;
  }

  function createResultsBlock() {
    const frame = document.createElement('iframe');
    // TODO: calculate height!
    // TODO: fix width, horyzontal scroll on google page
    frame.style = 'width: 454px; height: 450px; border: none;';
    frame.src = chrome.runtime.getURL('searchResults.html');
    return frame;
  }

  async function app() {
    const rightBlock = findOrCreateBlock();
    const resultsBlock = createResultsBlock();
    rightBlock.prepend(resultsBlock);
  }

  return app();
})();
