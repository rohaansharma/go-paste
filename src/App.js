import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import PasteForm from './components/PasteForm';
import PasteViewer from './components/PasteViewer';
import {Navbar, Container} from 'react-bootstrap';

function App() {
    return (
        <Router>
            <Navbar bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="/">Paste</Navbar.Brand>
                </Container>
            </Navbar>
            <Routes>
                <Route exact path="/" element={<PasteForm/>}/>
                <Route path="/paste/:id" element={<PasteViewer/>}/>
            </Routes>
        </Router>
    );
}

export default App;