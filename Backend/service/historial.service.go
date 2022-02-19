package service

import (
	m "sopes-practica1/models"
	historialRepository "sopes-practica1/repository"
)

func Create(his m.Historial) error {
	err := historialRepository.Create(his)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Historiales, error) {
	his, err := historialRepository.Read()

	if err != nil {
		return nil, err
	}
	return his, nil
}
