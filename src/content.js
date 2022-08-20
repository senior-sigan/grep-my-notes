(() => {
  function getGoogleQuery() {
    const params = new URLSearchParams(window.location.search);
    return params.get('q');
  }

  function app() {
    let rightBlock = document.getElementById('rhs');
    if (!rightBlock) {
      const searchPage = document.getElementById('rcnt');
      rightBlock = document.createElement('div');
      rightBlock.id = 'rhs';
      rightBlock.role = 'complementary';
      searchPage.append(rightBlock);
    }
    const searchResults = document.createElement('iframe');
    searchResults.dataset.query = getGoogleQuery();
    searchResults.style = 'width: 454px; height: 265px; border: none;';
    searchResults.src = chrome.runtime.getURL('searchResults.html');
    rightBlock.prepend(searchResults);
  }

  app();
})();
