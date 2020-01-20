package com.intertech.docker.Tour;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class TourService {

    @Autowired
    private TourRepository tourRepository;


    Tour save(Tour tour) {
        return tourRepository.save(tour);
    }
}
