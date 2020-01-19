package com.intertech.docker.Tour;

import lombok.*;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import java.util.Date;
import java.util.UUID;

@Entity
@EqualsAndHashCode(onlyExplicitlyIncluded = true)
@Builder
@AllArgsConstructor
public class Tour {

    @Column
    @Getter
    @Setter
    private String name;

    @Column(nullable = false, updatable = false)
    @CreationTimestamp
    @Getter
    private Date creationTimestamp;

    @Getter
    @Setter
    private String description;

    @Id
    @Getter
    @EqualsAndHashCode.Include
    private String id;

    @Column(nullable = false)
    @Getter
    @Setter
    private boolean shared = false;

    public Tour() {
        this.id = UUID.randomUUID().toString();
    }

}
