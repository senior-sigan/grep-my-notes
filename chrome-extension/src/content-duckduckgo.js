(async () => {
  function findOrCreateBlock() {
    console.log('BUM');
    const block = document.querySelector('section[data-area="sidebar"]');
    return block;
  }

  function createResultsBlock() {
    const frame = document.createElement('iframe');
    // TODO: calculate height!
    // TODO: fix width, horyzontal scroll on google page
    frame.style = 'width: 450px; height: 900px; border: none;';
    frame.src = chrome.runtime.getURL('searchResults.html');
    return frame;
  }

  async function app() {
    const rightBlock = findOrCreateBlock();
    const resultsBlock = createResultsBlock();

    console.log(rightBlock);
    rightBlock.prepend(resultsBlock);
  }

  return app();
})();
