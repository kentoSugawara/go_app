import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import ListComponent from './ListComponent';
import LoginComponent from './LoginComponent';

function AppComponent() {
  return (
    <Router>
      <Routes>
        <Route path="/list" element={<ListComponent />} />
        <Route path="/login" element={<LoginComponent />} />
      </Routes>
    </Router>
  );
}

export default AppComponent;